"use client";

import DynamicTitle from "@/components/DynamicTitle";
import React, { useState } from "react";
import { UserProps } from "@/types/UserType";
import { useRouter } from "next/navigation";

export default function UpdateForm({ user }: { user: UserProps }) {
  const router = useRouter();

  const [formData, setFormData] = useState({
    name: user.name,
    email: user.email,
  });

  const { name, email } = formData;

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      if (!name || !email) {
        console.log("Please fill all fields");
        return;
      }

      const res = await fetch(`http://0.0.0.0:8080/users/${user.id}`, {
        cache: "no-store",
        method: "PATCH",
        body: JSON.stringify(formData),
        headers: {
          "Content-Type": "application/json",
        },
      });

      const data = await res.json();

      if (!res.ok) {
        console.log(data.error);
      }

      if (res.status === 200) {
        console.log(data.message);
        router.push(`/users/${user.id}`);
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <DynamicTitle title="USER GUNCELLE PAGE" />
      <form
        className="flex flex-col gap-2 w-1/2 mx-auto text-black"
        onSubmit={handleSubmit}
      >
        <input
          type="text"
          placeholder="name"
          name="name"
          value={name}
          onChange={handleChange}
        />
        <input
          type="email"
          placeholder="email"
          name="email"
          value={email}
          onChange={handleChange}
        />
        <button type="submit" className="bg-yellow-600 text-white">
          GUNCELLE
        </button>
      </form>
    </div>
  );
}
