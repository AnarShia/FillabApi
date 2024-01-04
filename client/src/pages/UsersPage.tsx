import useSWR from "swr";
import { Box, Button, List, MantineProvider, TextInput, ThemeIcon } from "@mantine/core";
import { useState } from "react";
import Layout from "../layouts/Layout";

export const ENDPOINT = "http://localhost:4000";

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

interface User {
  id: number
  username: string;
}


function UsersPage() {
  const { data, mutate } = useSWR<User[]>("users", fetcher);
  const [newUser, setNewUser] = useState({ username: '' });
  const [editUser, setEditUser] = useState<User | null>(null);

  const createUser = async (e: React.FormEvent) => {
    e.preventDefault();

    const response = await fetch(`${ENDPOINT}/users`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newUser),
    });

    if (!response.ok) {
      throw new Error('Error creating user');
    }

    const createdUser = await response.json();
    mutate([...(data!), createdUser], false);
    setNewUser({ username: '' });
  };

  const updateUser = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!editUser) return;

    const response = await fetch(`${ENDPOINT}/users/${editUser.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(editUser),
    });

    if (!response.ok) {
      throw new Error('Error updating user');
    }

    const updatedUser = await response.json();
    mutate(data?.map((user) => user.id === updatedUser.id ? updatedUser : user), false);
    setEditUser(null);
  };

  const deleteUser = async (id: number) => {
    const response = await fetch(`${ENDPOINT}/users/${id}`, {
      method: 'DELETE',
    });

    if (!response.ok) {
      throw new Error('Error deleting user');
    }

    mutate(data?.filter((user) => user.id !== id), false);
  };

  return (
    <Layout>
      <MantineProvider>
        <Box className="users-page-box">
          <List>
            {data?.map((user) => (
              <List.Item key={user.id} className="list-item" icon={<ThemeIcon variant="dot" color="blue" />}>
                <span className="list-item-span">{user.username ?? 'Unknown'}</span>
                <button className="list-item-btn" onClick={() => deleteUser(user.id)}>Delete</button>
                <button className="list-item-btn" onClick={() => setEditUser(user)}>Edit</button>
              </List.Item>
            ))}
          </List>
          <form className="users-page-form" onSubmit={editUser ? updateUser : createUser}>
            <TextInput className="input"
              value={editUser ? editUser.username : newUser.username}
              onChange={(e) => editUser ? setEditUser({ ...editUser, username: e.target.value }) : setNewUser({ ...newUser, username: e.target.value })}
              placeholder="User Name"
              required
            />
            <Button type="submit">{editUser ? 'Update User' : 'Create User'}</Button>
          </form>
        </Box>
      </MantineProvider>
    </Layout>
  );
}

export default UsersPage;