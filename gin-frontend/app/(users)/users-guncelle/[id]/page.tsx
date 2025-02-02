import EditForm from "@/components/user/edit-form";
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
      <EditForm user={user} />
    </div>
  );
}
