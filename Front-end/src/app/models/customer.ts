export interface Customer {
  id: number;
  firstName: string;
  lastName: string;
  dateOfBirth?: Date;
  address: string;
  delFlag: boolean;
  age?: number;
}
