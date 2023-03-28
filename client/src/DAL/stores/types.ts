export enum DataObjectState {
  NotStarted,
  Fetching,
  Adding,
  Removing,
  Succeeded,
  Failed,
}

export type DataObject<T> = {
  data: T
  state: DataObjectState
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  error?: any
}
