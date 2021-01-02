import { DateTime } from "luxon";

export default function (isoDateTime) {
  return DateTime.fromISO(isoDateTime).toLocaleString(DateTime.DATETIME_MED);
}
