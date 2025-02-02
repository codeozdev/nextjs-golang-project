import UpdateForm from "@/components/user/update-form";
import { getUser } from "@/fetchData/get-user";

export default async function UsersGuncelle({
  params,
}: {
  params: { id: string };
}) {
  const { id } = await params;
  const { user } = await getUser(id);

  return (
    <div>
      <UpdateForm user={user} />
    </div>
  );
}
