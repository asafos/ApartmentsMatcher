export enum DataObjectState {
    NotStarted,
    Fetching,
    Adding,
    Removing,
    Succeeded,
    Failed
}

export type DataObject<T> = {
    data: T,
    state: DataObjectState,
    error?: any
}
