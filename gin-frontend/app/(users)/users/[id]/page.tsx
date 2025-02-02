import DynamicTitle from "@/components/DynamicTitle";
import { getUser } from "@/fetchData/get-user";
import Custom404 from "@/app/(users)/users/[id]/404";
import Link from "next/link";

export { getUser } from "@/fetchData/get-user";

export default async function UserPage({ params }: { params: { id: string } }) {
  const { id } = await params;
  const { user } = await getUser(id);

  if (!user) return <Custom404 />;

  return (
    <div>
      <DynamicTitle title="USER PAGE" />
      <Link href={`/users-guncelle/${user.id}`} className="bg-yellow-600 px-4">
        GUNCELLE
      </Link>
      <div className="flex gap-2 justify-between font-bold text-3xl">
        <h1>{user.name}</h1>
        <p>{user.email}</p>
      </div>
    </div>
  );
}
