import dayjs from "dayjs";
import "dayjs/locale/zh-cn";
import relativeTime from "dayjs/plugin/relativeTime";
import customParseFormat from "dayjs/plugin/customParseFormat";

// 配置 dayjs
dayjs.extend(relativeTime);
dayjs.extend(customParseFormat);
dayjs.locale("zh-cn");

/**
 * 格式化日期为中文格式
 * @param date 日期字符串或 Date 对象
 * @param format 格式化模板，默认为 "YYYY年MM月DD日 HH:mm"
 * @returns 格式化后的日期字符串
 */
export const formatDate = (
  date: string | Date,
  format = "YYYY年MM月DD日 HH:mm"
): string => {
  return dayjs(date).format(format);
};

/**
 * 格式化日期为简短格式
 * @param date 日期字符串或 Date 对象
 * @returns 格式化后的日期字符串 (YYYY-MM-DD HH:mm)
 */
export const formatDateShort = (date: string | Date): string => {
  return dayjs(date).format("YYYY-MM-DD HH:mm");
};

/**
 * 格式化日期为表格显示格式
 * @param date 日期字符串或 Date 对象
 * @returns 格式化后的日期字符串 (MM/DD HH:mm)
 */
export const formatDateTable = (date: string | Date): string => {
  return dayjs(date).format("YYYY-MM-DD HH:mm");
};

/**
 * 格式化日期为详细格式
 * @param date 日期字符串或 Date 对象
 * @returns 格式化后的日期字符串 (YYYY年MM月DD日 HH:mm:ss)
 */
export const formatDateDetail = (date: string | Date): string => {
  return dayjs(date).format("YYYY年MM月DD日 HH:mm:ss");
};

/**
 * 获取相对时间
 * @param date 日期字符串或 Date 对象
 * @returns 相对时间字符串 (如: "2小时前", "3天前")
 */
export const fromNow = (date: string | Date): string => {
  return dayjs(date).fromNow();
};

/**
 * 获取距离现在的时间
 * @param date 日期字符串或 Date 对象
 * @returns 距离时间字符串 (如: "还有2小时", "还有3天")
 */
export const toNow = (date: string | Date): string => {
  return dayjs(date).toNow();
};

/**
 * 判断日期是否为今天
 * @param date 日期字符串或 Date 对象
 * @returns 是否为今天
 */
export const isToday = (date: string | Date): boolean => {
  return dayjs(date).isSame(dayjs(), "day");
};

/**
 * 判断日期是否为昨天
 * @param date 日期字符串或 Date 对象
 * @returns 是否为昨天
 */
export const isYesterday = (date: string | Date): boolean => {
  return dayjs(date).isSame(dayjs().subtract(1, "day"), "day");
};

/**
 * 判断日期是否为本周
 * @param date 日期字符串或 Date 对象
 * @returns 是否为本周
 */
export const isThisWeek = (date: string | Date): boolean => {
  return dayjs(date).isSame(dayjs(), "week");
};

/**
 * 智能格式化日期
 * 今天显示时间，昨天显示"昨天 HH:mm"，本周显示"周X HH:mm"，其他显示完整日期
 * @param date 日期字符串或 Date 对象
 * @returns 智能格式化后的日期字符串
 */
export const formatDateSmart = (date: string | Date): string => {
  const target = dayjs(date);
  const now = dayjs();

  if (target.isSame(now, "day")) {
    return target.format("HH:mm");
  } else if (target.isSame(now.subtract(1, "day"), "day")) {
    return `昨天 ${target.format("HH:mm")}`;
  } else if (target.isSame(now, "week")) {
    const weekdays = ["周日", "周一", "周二", "周三", "周四", "周五", "周六"];
    return `${weekdays[target.day()]} ${target.format("HH:mm")}`;
  } else if (target.isSame(now, "year")) {
    return target.format("MM月DD日 HH:mm");
  } else {
    return target.format("YYYY年MM月DD日 HH:mm");
  }
};

/**
 * 获取当前时间戳
 * @returns 当前时间戳
 */
export const now = (): number => {
  return dayjs().valueOf();
};

/**
 * 获取当前日期字符串
 * @param format 格式化模板，默认为 "YYYY-MM-DD HH:mm:ss"
 * @returns 当前日期字符串
 */
export const nowString = (format = "YYYY-MM-DD HH:mm:ss"): string => {
  return dayjs().format(format);
};
