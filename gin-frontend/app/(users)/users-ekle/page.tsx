"use client";

import DynamicTitle from "@/components/DynamicTitle";

import React, { useState } from "react";

export default function UsersEkle() {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
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

      const res = await fetch("http://0.0.0.0:8080/users", {
        cache: "no-store",
        method: "POST",
        body: JSON.stringify(formData),
        headers: {
          "Content-Type": "application/json",
        },
      });

      const data = await res.json();
      console.log(data);

      if (!res.ok) {
        console.log("Error");
      }

      if (res.status === 200) {
        console.log(data.message);
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <DynamicTitle title="USERS EKLE PAGE" />
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
        <button type="submit" className="bg-green-500">
          EKLE
        </button>
      </form>
    </div>
  );
}
