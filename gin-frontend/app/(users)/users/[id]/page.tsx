import DynamicTitle from "@/components/DynamicTitle";
import { getUser } from "@/fetchData/get-user";

export { getUser } from "@/fetchData/get-user";

export default async function UserPage({ params }: { params: { id: string } }) {
  const { id } = await params;
  const { user } = await getUser(id);

  return (
    <div>
      <DynamicTitle title="USER PAGE" />
      <div className="flex gap-2 justify-between font-bold text-3xl">
        <h1>{user.name}</h1>
        <p>{user.email}</p>
      </div>
    </div>
  );
}
