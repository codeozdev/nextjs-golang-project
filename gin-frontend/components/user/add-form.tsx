"use client";

import DynamicTitle from "@/components/DynamicTitle";
import React, { useState } from "react";
import { useRouter } from "next/navigation";

export default function AddForm() {
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

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
    setError(null); // Form gönderildiğinde hata mesajını temizle

    try {
      if (!name || !email) {
        setError("Please fill all fields");
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

      if (!res.ok) {
        console.log(data.error);
        setError(data.error);
      }

      if (res.status === 400) {
        setError(data.error);
        console.log(data.error);
        return;
      }

      if (res.status === 201) {
        console.log(data.message);
        router.push(`/users`);
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <DynamicTitle title="USER EKLE PAGE" />
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
      {/* Hata mesajını göster */}
      {error && (
        <div className="mt-4 text-center text-red-500">
          <p>{error}</p>
        </div>
      )}
    </div>
  );
}
