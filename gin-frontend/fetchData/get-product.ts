const URL = process.env.NEXT_PUBLIC_API_URL;

export const getAllProducts = async () => {
  const res = await fetch(`${URL}/products`);
  return await res.json();
};
