import {
  createMandatoryNumberCriterionOption,
  createStringCriterionOption,
  NullNumberCriterionOption,
  createDateCriterionOption,
  createMandatoryTimestampCriterionOption,
  createPathCriterionOption,
} from "./criteria/criterion";
import { PerformerFavoriteCriterionOption } from "./criteria/favorite";
import { GalleryIsMissingCriterionOption } from "./criteria/is-missing";
import { OrganizedCriterionOption } from "./criteria/organized";
import { HasChaptersCriterionOption } from "./criteria/has-chapters";
import { PerformersCriterionOption } from "./criteria/performers";
import { AverageResolutionCriterionOption } from "./criteria/resolution";
import { StudiosCriterionOption } from "./criteria/studios";
import {
  PerformerTagsCriterionOption,
  TagsCriterionOption,
} from "./criteria/tags";
import { ListFilterOptions, MediaSortByOptions } from "./filter-options";
import { DisplayMode } from "./types";

// 原版的默认排序是路径，但这样顺序过于固定不利于查看新数据，改为按照日期排序
const defaultSortBy = "date";

// 排序窗口可选的排序字段
const sortByOptions = ["date", ...MediaSortByOptions]
  .map(ListFilterOptions.createSortBy)
  .concat([
    {
      messageID: "image_count",
      value: "images_count",
    },
    {
      messageID: "zip_file_count",
      value: "file_count",
    },
  ]);

// 显示方式（方格、列表、照片墙）
const displayModeOptions = [
  DisplayMode.Grid,
  DisplayMode.List,
  DisplayMode.Wall,
];

const criterionOptions = [
  createStringCriterionOption("title"),
  createStringCriterionOption("details"),
  createPathCriterionOption("path"),
  createStringCriterionOption(
    "galleryChecksum",
    "media_info.checksum",
    "checksum"
  ),
  new NullNumberCriterionOption("rating", "rating100"),
  OrganizedCriterionOption,
  AverageResolutionCriterionOption,
  GalleryIsMissingCriterionOption,
  TagsCriterionOption,
  HasChaptersCriterionOption,
  createStringCriterionOption("tag_count"),
  PerformerTagsCriterionOption,
  PerformersCriterionOption,
  createStringCriterionOption("performer_count"),
  createMandatoryNumberCriterionOption("performer_age"),
  PerformerFavoriteCriterionOption,
  createStringCriterionOption("image_count"),
  StudiosCriterionOption,
  createStringCriterionOption("url"),
  createMandatoryNumberCriterionOption("file_count", "zip_file_count"),
  createDateCriterionOption("date"),
  createMandatoryTimestampCriterionOption("created_at"),
  createMandatoryTimestampCriterionOption("updated_at"),
];

export const GalleryListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
