"use client";

import { useParams } from "next/navigation";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function Custom404() {
  const { id } = useParams();
  const router = useRouter();

  const [countdown, setCountdown] = useState(5); // Geri sayım süresi (5 saniye)

  // kullaniciyi bulamazsa 5 saniye sonra users sayfasina yönlendirir
  useEffect(() => {
    // 1 saniyede bir geri sayımı azalt
    const timer = setInterval(() => {
      setCountdown((prev) => prev - 1);
    }, 1000);

    // 5 saniye sonra yönlendirme yap
    const redirectTimer = setTimeout(() => {
      router.push("/users");
    }, 5000);

    // Component unmount olduğunda timer'ları temizle
    return () => {
      clearInterval(timer);
      clearTimeout(redirectTimer);
    };
  }, [router]);

  return (
    <div className="flex flex-col items-center justify-center text-3xl font-bold">
      <h1>404</h1>
      <p>{id} kullanıcı bulunamadı</p>
      <p className="text-base">
        {countdown} saniye sonra ana sayfaya yönlendirileceksiniz...
      </p>
    </div>
  );
}
