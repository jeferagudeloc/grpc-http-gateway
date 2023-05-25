import { UsersResponseMock } from "../../__mocks__/usersMock";

export interface User {
  id: string;
  name: string;
  lastname: string;
  email: string;
  status: string;
  role: Role;
}

export interface Role {
  name: string;
  permissions: string[];
}

export interface UserState {
  data: User[];
}

export const UserStateDefault: UserState = {
  data: UsersResponseMock,
};
