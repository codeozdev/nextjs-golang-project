import UsersComponent from "@/components/UsersComponent";
import { UserProps } from "@/types/UserType";
import { getAllUser } from "@/fetchData/get-user";
import DynamicTitle from "@/components/DynamicTitle";
import Link from "next/link";

export default async function UsersPage() {
  const { message, users } = await getAllUser();
  console.log(message);
  return (
    <div>
      <DynamicTitle title="USERS PAGE" />
      <div className="flex justify-between px-2  text-xl font-bold">
        <Link href="/users-ekle" className="bg-blue-500 px-4">
          EKLE
        </Link>
        <Link href="/users-sil" className="bg-red-500 px-4">
          SIL
        </Link>
      </div>
      {users.map((user: UserProps) => (
        <UsersComponent key={user.id} user={user} />
      ))}
    </div>
  );
}
