"use client";
import { Shell } from "lucide-react";
import { LoginForm } from "./LoginForm";
import { SignupForm } from "./SignupForm";
import { useState } from "react";

export default function LoginSignUp() {
  const [toggleActions, setToggleActions] = useState(false);

  return (
    <div className="grid min-h-screen lg:grid-cols-2">
      <div className="flex flex-col gap-4 p-6 md:p-10 bg-white">
        <div className="flex justify-center gap-2 md:justify-start">
          <a href="#" className="flex items-center gap-2 font-medium">
            <div className="bg-blue-600 text-white flex w-6 h-6 items-center justify-center rounded-md">
              <Shell className="w-4 h-4" />
            </div>
            <p className="font-semibold">Coil Inc.</p>
          </a>
        </div>
        <div className="flex flex-1 items-center justify-center">
          <div className="w-full max-w-xs">
            {toggleActions == false ? (
              <LoginForm setToggleActions={setToggleActions} />
            ) : (
              <SignupForm setToggleActions={setToggleActions} />
            )}
          </div>
        </div>
      </div>
      <div className="bg-gray-100 relative hidden lg:block">
        <img
          src="/login.png"
          alt="Image"
          className="absolute inset-0 h-full w-full object-cover"
        />
      </div>
    </div>
  );
}
