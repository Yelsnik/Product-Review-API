"use client"; // Error boundaries must be Client Components

import { useEffect } from "react";

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    // Log the error to an error reporting service
    console.error(error, error.name, error.cause);
  }, [error]);

  return (
    <div className="py-10">
      <div className="grid mt-40 justify-center">
        <h2 className="text-bold text-4xl ">Something went wrong!</h2>
        <div>
          <button
            onClick={
              // Attempt to recover by trying to re-render the segment
              () => reset()
            }
            className="duration-700 ease-in-out hover:-translate-y-1 hover:scale-105 mt-2 border rounded bg-stone-600 hover:bg-stone-800 border-transparent px-2 py-2"
          >
            Try again
          </button>
        </div>
      </div>
    </div>
  );
}
