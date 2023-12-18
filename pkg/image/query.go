package image

import (
	"context"
	"strconv"

	"github.com/stashapp/stash/pkg/models"
)

type Queryer interface {
	Query(ctx context.Context, options models.ImageQueryOptions) (*models.ImageQueryResult, error)
}

type CountQueryer interface {
	QueryCount(ctx context.Context, imageFilter *models.ImageFilterType, findFilter *models.FindFilterType) (int, error)
}

// QueryOptions returns a ImageQueryResult populated with the provided filters.
func QueryOptions(imageFilter *models.ImageFilterType, findFilter *models.FindFilterType, count bool) models.ImageQueryOptions {
	return models.ImageQueryOptions{
		QueryOptions: models.QueryOptions{
			FindFilter: findFilter,
			Count:      count,
		},
		ImageFilter: imageFilter,
	}
}

// Query queries for images using the provided filters.
func Query(ctx context.Context, qb Queryer, imageFilter *models.ImageFilterType, findFilter *models.FindFilterType) ([]*models.Image, error) {
	result, err := qb.Query(ctx, QueryOptions(imageFilter, findFilter, false))
	if err != nil {
		return nil, err
	}

	images, err := result.Resolve(ctx)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func CountByPerformerID(ctx context.Context, r CountQueryer, id int) (int, error) {
	filter := &models.ImageFilterType{
		Performers: &models.MultiCriterionInput{
			Value:    []string{strconv.Itoa(id)},
			Modifier: models.CriterionModifierIncludes,
		},
	}

	return r.QueryCount(ctx, filter, nil)
}

func CountByStudioID(ctx context.Context, r CountQueryer, id int, depth *int) (int, error) {
	filter := &models.ImageFilterType{
		Studios: &models.HierarchicalMultiCriterionInput{
			Value:    []string{strconv.Itoa(id)},
			Modifier: models.CriterionModifierIncludes,
			Depth:    depth,
		},
	}

	return r.QueryCount(ctx, filter, nil)
}

func CountByTagID(ctx context.Context, r CountQueryer, id int, depth *int) (int, error) {
	filter := &models.ImageFilterType{
		Tags: &models.HierarchicalMultiCriterionInput{
			Value:    []string{strconv.Itoa(id)},
			Modifier: models.CriterionModifierIncludes,
			Depth:    depth,
		},
	}

	return r.QueryCount(ctx, filter, nil)
}

func FindByGalleryID(ctx context.Context, r Queryer, galleryID int, sortBy string, sortDir models.SortDirectionEnum) ([]*models.Image, error) {
	perPage := -1

	findFilter := models.FindFilterType{
		PerPage: &perPage,
	}

	if sortBy != "" {
		findFilter.Sort = &sortBy
	}

	if sortDir.IsValid() {
		findFilter.Direction = &sortDir
	}

	return Query(ctx, r, &models.ImageFilterType{
		Galleries: &models.MultiCriterionInput{
			Value:    []string{strconv.Itoa(galleryID)},
			Modifier: models.CriterionModifierIncludes,
		},
	}, &findFilter)
}

func FindGalleryCover(ctx context.Context, r Queryer, galleryID int, galleryCoverRegex string) (*models.Image, error) {
	const useCoverJpg = true
	img, err := findGalleryCover(ctx, r, galleryID, useCoverJpg, galleryCoverRegex)
	if err != nil {
		return nil, err
	}

	if img != nil {
		return img, nil
	}

	// return the first image in the gallery
	return findGalleryCover(ctx, r, galleryID, !useCoverJpg, galleryCoverRegex)
}

func findGalleryCover(ctx context.Context, r Queryer, galleryID int, useCoverJpg bool, galleryCoverRegex string) (*models.Image, error) {
	// try to find cover.jpg in the gallery
	perPage := 1
	// 原版的默认排序是路径，但是按照路径排序会导致不同图集里的图片都按照主文件的路径排序，这就导致有的图集里这个图片永远在最前面而有的则永远在最后面
	// 实际上这个问题是系统设计的硬伤无法通过简单的修补解决，这里利用图片很少改名的特性用标题作为排序标准
	sortBy := "title"
	sortDir := models.SortDirectionEnumAsc

	findFilter := models.FindFilterType{
		PerPage:   &perPage,
		Sort:      &sortBy,
		Direction: &sortDir,
	}

	imageFilter := &models.ImageFilterType{
		Galleries: &models.MultiCriterionInput{
			Value:    []string{strconv.Itoa(galleryID)},
			Modifier: models.CriterionModifierIncludes,
		},
	}

	if useCoverJpg {
		imageFilter.Path = &models.StringCriterionInput{
			Value:    "(?i)" + galleryCoverRegex,
			Modifier: models.CriterionModifierMatchesRegex,
		}
	}

	imgs, err := Query(ctx, r, imageFilter, &findFilter)
	if err != nil {
		return nil, err
	}

	if len(imgs) > 0 {
		return imgs[0], nil
	}

	return nil, nil
}
