const URL = process.env.NEXT_PUBLIC_API_URL;

export const getAllUser = async () => {
  const res = await fetch(`${URL}/users`);
  return await res.json();
};

export const getUser = async (id: string) => {
  const res = await fetch(`${URL}/users/${id}`);
  return await res.json();
};
