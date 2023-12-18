import {
  createMandatoryNumberCriterionOption,
  createStringCriterionOption,
  NullNumberCriterionOption,
  createDateCriterionOption,
  createMandatoryTimestampCriterionOption,
} from "./criteria/criterion";
import { MovieIsMissingCriterionOption } from "./criteria/is-missing";
import { StudiosCriterionOption } from "./criteria/studios";
import { PerformersCriterionOption } from "./criteria/performers";
import { ListFilterOptions } from "./filter-options";
import { DisplayMode } from "./types";

// 原版的默认排序是标题，但这样顺序过于固定不利于查看新数据，改为按照日期排序
const defaultSortBy = "date";

const sortByOptions = ["name", "random", "date", "duration", "rating"]
  .map(ListFilterOptions.createSortBy)
  .concat([
    {
      messageID: "scene_count",
      value: "scenes_count",
    },
  ]);
const displayModeOptions = [DisplayMode.Grid];
const criterionOptions = [
  StudiosCriterionOption,
  MovieIsMissingCriterionOption,
  createStringCriterionOption("url"),
  createStringCriterionOption("name"),
  createStringCriterionOption("director"),
  createStringCriterionOption("synopsis"),
  createMandatoryNumberCriterionOption("duration"),
  new NullNumberCriterionOption("rating", "rating100"),
  PerformersCriterionOption,
  createDateCriterionOption("date"),
  createMandatoryTimestampCriterionOption("created_at"),
  createMandatoryTimestampCriterionOption("updated_at"),
];

export const MovieListFilterOptions = new ListFilterOptions(
  defaultSortBy,
  sortByOptions,
  displayModeOptions,
  criterionOptions
);
