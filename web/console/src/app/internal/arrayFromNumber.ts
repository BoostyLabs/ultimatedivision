// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Creates objects array from number. */
export const arrayFromNumber = (
  amount: number
): {
  id: number;
  [key: string]: any;
}[] => {
  const array = [];
  for (let index = 0; index < amount; index++) {
    array.push({ id: index });
  }
  return array;
};
