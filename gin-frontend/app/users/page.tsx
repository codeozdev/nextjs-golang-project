import UsersComponent from "@/components/UsersComponent";
import { User } from "@/types/UserType";
import { getAllUser } from "@/fetchData/get-user";
import DynamicTitle from "@/components/DynamicTitle";

export default async function UsersPage() {
  const { message, users } = await getAllUser();
  console.log(message);
  return (
    <div>
      <DynamicTitle title="USERS PAGE" />
      {users.map((user: User) => (
        <UsersComponent key={user.id} user={user} />
      ))}
    </div>
  );
}
