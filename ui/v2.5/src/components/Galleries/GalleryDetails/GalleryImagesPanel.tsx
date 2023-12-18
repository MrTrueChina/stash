import React from "react";
import * as GQL from "src/core/generated-graphql";
import { GalleriesCriterion } from "src/models/list-filter/criteria/galleries";
import { ListFilterModel } from "src/models/list-filter/filter";
import { ImageList } from "src/components/Images/ImageList";
import { mutateRemoveGalleryImages } from "src/core/StashService";
import {
  showWhenSelected,
  PersistanceLevel,
} from "src/components/List/ItemList";
import { useToast } from "src/hooks/Toast";
import { useIntl } from "react-intl";
import { faMinus } from "@fortawesome/free-solid-svg-icons";
import { galleryTitle } from "src/core/galleries";

interface IGalleryDetailsProps {
  active: boolean;
  gallery: GQL.GalleryDataFragment;
}

// 显示一个图集内部的图片的界面
export const GalleryImagesPanel: React.FC<IGalleryDetailsProps> = ({
  active,
  gallery,
}) => {
  const intl = useIntl();
  const Toast = useToast();

  // 似乎是对事件的监听，Stash 里面对此的描述是 "Hook(钩子/挂钩)"
  function filterHook(filter: ListFilterModel) {
    const galleryValue = {
      id: gallery.id!,
      label: galleryTitle(gallery),
    };
    // if galleries is already present, then we modify it, otherwise add
    let galleryCriterion = filter.criteria.find((c) => {
      return c.criterionOption.type === "galleries";
    }) as GalleriesCriterion;

    if (
      galleryCriterion &&
      (galleryCriterion.modifier === GQL.CriterionModifier.IncludesAll ||
        galleryCriterion.modifier === GQL.CriterionModifier.Includes)
    ) {
      // add the gallery if not present
      if (
        !galleryCriterion.value.find((p) => {
          return p.id === gallery.id;
        })
      ) {
        galleryCriterion.value.push(galleryValue);
      }

      galleryCriterion.modifier = GQL.CriterionModifier.IncludesAll;
    } else {
      // overwrite
      galleryCriterion = new GalleriesCriterion();
      galleryCriterion.value = [galleryValue];
      filter.criteria.push(galleryCriterion);
    }

    return filter;
  }

  async function removeImages(
    result: GQL.FindImagesQueryResult,
    filter: ListFilterModel,
    selectedIds: Set<string>
  ) {
    try {
      await mutateRemoveGalleryImages({
        gallery_id: gallery.id!,
        image_ids: Array.from(selectedIds.values()),
      });

      Toast.success({
        content: intl.formatMessage(
          { id: "toast.removed_entity" },
          {
            count: selectedIds.size,
            singularEntity: intl.formatMessage({ id: "image" }),
            pluralEntity: intl.formatMessage({ id: "images" }),
          }
        ),
      });
    } catch (e) {
      Toast.error(e);
    }
  }

  const otherOperations = [
    {
      text: intl.formatMessage({ id: "actions.remove_from_gallery" }),
      onClick: removeImages,
      isDisplayed: showWhenSelected,
      postRefetch: true,
      icon: faMinus,
      buttonVariant: "danger",
    },
  ];

  // 最后返回了一个 ImageList，从名字上看是图片列表，可能就是图片显示的那部分，而这个组件估计是包含了外部的各种操作按钮
  return (
    <ImageList
      filterHook={filterHook}
      alterQuery={active}
      extraOperations={otherOperations}
      persistState={PersistanceLevel.VIEW}
      persistanceKey="galleryimages"
      chapters={gallery.chapters}
    />
  );
};
