export enum DataObjectState {
    NotStarted,
    InProgress,
    Succeeded,
    Failed
}

export type DataObject<T> = {
    data: T,
    state: DataObjectState,
    error?: any
}
