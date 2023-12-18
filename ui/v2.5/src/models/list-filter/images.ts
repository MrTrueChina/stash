import {
  createMandatoryNumberCriterionOption,
  createMandatoryStringCriterionOption,
  createStringCriterionOption,
  NullNumberCriterionOption,
  createMandatoryTimestampCriterionOption,
  createDateCriterionOption,
  createPathCriterionOption,
} from "./criteria/criterion";
import { PerformerFavoriteCriterionOption } from "./criteria/favorite";
import { ImageIsMissingCriterionOption } from "./criteria/is-missing";
import { OrganizedCriterionOption } from "./criteria/organized";
import { PerformersCriterionOption } from "./criteria/performers";
import { ResolutionCriterionOption } from "./criteria/resolution";
import { StudiosCriterionOption } from "./criteria/studios";
import {
  PerformerTagsCriterionOption,
  TagsCriterionOption,
} from "./criteria/tags";
import { ListFilterOptions, MediaSortByOptions } from "./filter-options";
import { DisplayMode } from "./types";

// 原版的默认排序是路径，但是按照路径排序会导致不同图集里的图片都按照主文件的路径排序，这就导致有的图集里这个图片永远在最前面而有的则永远在最后面
// 实际上这个问题是系统设计的硬伤无法通过简单的修补解决，这里利用图片很少改名的特性用标题作为排序标准
const defaultSortBy = "title";

const sortByOptions = [
  "o_counter",
  "filesize",
  "file_count",
  "date",
  ...MediaSortByOptions,
].map(ListFilterOptions.createSortBy);

const displayModeOptions = [DisplayMode.Grid, DisplayMode.Wall];
const criterionOptions = [
  createStringCriterionOption("title"),
  createMandatoryStringCriterionOption("checksum", "media_info.checksum"),
  createPathCriterionOption("path"),
  OrganizedCriterionOption,
  createMandatoryNumberCriterionOption("o_counter"),
  ResolutionCriterionOption,
  ImageIsMissingCriterionOption,
  TagsCriterionOption,
  new NullNumberCriterionOption("rating", "rating100"),
  createMandatoryNumberCriterionOption("tag_count"),
  PerformerTagsCriterionOption,
  PerformersCriterionOption,
  createMandatoryNumberCriterionOption("performer_count"),
  PerformerFavoriteCriterionOption,
  StudiosCriterionOption,
  createStringCriterionOption("url"),
  createDateCriterionOption("date"),
  createMandatoryNumberCriterionOption("file_count"),
  createMandatoryTimestampCriterionOption("created_at"),
  createMandatoryTimestampCriterionOption("updated_at"),
];
export const ImageListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
