enum Location {
  LevTLV = 'LevTLV',
  OldNorth = 'OldNorth',
  NewNorth = 'NewNorth',
  Sarona = 'Sarona',
  NeveTzedek = 'NeveTzedek',
  NeveShaanan = 'NeveShaanan',
  Florentin = 'Florentin',
  RamatAviv = 'RamatAviv',
}

const locationLabelsMap: Record<Location, string> = {
  [Location.LevTLV]: 'Lev TLV',
  [Location.OldNorth]: 'Old North',
  [Location.NewNorth]: 'New North',
  [Location.Sarona]: 'Sarona',
  [Location.NeveTzedek]: 'Neve Tzedek',
  [Location.NeveShaanan]: 'Neve Shaanan',
  [Location.Florentin]: 'Florentin',
  [Location.RamatAviv]: 'Ramat Aviv ',
}

export { Location, locationLabelsMap }
