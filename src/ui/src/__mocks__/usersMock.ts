import { User } from "../store/states/userState";

export const UsersResponseMock: User[] = [
    {
      id: "1",
      name: "John",
      lastname: "Doe",
      email: "john.doe@example.com",
      status: "active",
      role: {
        name: "admin",
        permissions: ["read", "write", "delete"],
      },
    },
    {
      id: "2",
      name: "Jane",
      lastname: "Smith",
      email: "jane.smith@example.com",
      status: "active",
      role: {
        name: "user",
        permissions: ["read"],
      },
    },
    {
      id: "3",
      name: "Alice",
      lastname: "Johnson",
      email: "alice.johnson@example.com",
      status: "inactive",
      role: {
        name: "user",
        permissions: ["read"],
      },
    },
  ];
  