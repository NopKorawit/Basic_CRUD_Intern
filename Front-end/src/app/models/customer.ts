export interface Customer {
  id: number;
  firstName: string;
  lastName: string;
  dateOfBirth: number;
  address: string;
  delFlag: boolean;
  age?: number;
}
