export const getAllUser = async () => {
  const res = await fetch("http://0.0.0.0:8080/users");
  return await res.json();
};

export const getUser = async (id: string) => {
  const res = await fetch(`http://0.0.0.0:8080/users/${id}`);
  return await res.json();
};
