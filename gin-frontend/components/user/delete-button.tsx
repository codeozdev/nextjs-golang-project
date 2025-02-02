"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

export default function DeleteButton({ id }: { id: string }) {
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleDelete = async () => {
    try {
      const res = await fetch(`http://0.0.0.0:8080/users/${id}`, {
        method: "DELETE",
      });

      const data = await res.json();

      if (!res.ok) {
        setError(data.error); // Backend'den gelen hatayı state'e kaydet
        return;
      }

      if (res.status === 400) {
        setError(data.error);
        return;
      }

      if (res.status === 404) {
        setError(data.error);
        return;
      }

      if (res.status === 200) {
        console.log(data.message);
        router.push(`/users`);
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      {error && (
        <div className="mt-4 text-center text-red-500">
          <p>{error}</p>
        </div>
      )}
      <button onClick={handleDelete} className="bg-red-500 px-4">
        Sil
      </button>
    </div>
  );
}
