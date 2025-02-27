import Image from "next/image";

import Link from "next/link";
import clsx from "clsx";

export default async function Home() {
  const api = await fetch("http://localhost:4000/v1/reviews?page=1&country=us");
  const data = await api.json();
  const products = data.data;

  return (
    <div>
      <nav className="border border-gray-800 border-transparent fixed top-0 left-0 right-0 z-50 bg-opacity-75 backdrop-filter backdrop-blur-lg items-center px-6 sm:px-10 lg:px-40 py-5 text-white bg-transparent">
        <div className="flex justify-between">
          <div>Product Review</div>
          <Link href={"leaderboard"}>Leaderboard</Link>
        </div>
      </nav>

      <div className="mt-20 px-5 sm:px-10 md:px-20 lg:px-40">
        <div className="max-w-xl mx-auto">
          <form className="flex flex-col sm:flex-row items-center gap-3 w-full">
            <input
              className="w-full sm:w-72 py-2 px-4 font-semibold text-black outline-none bg-gray-100 rounded focus:ring-2 focus:ring-green-600"
              type="text"
              placeholder="Search for a product..."
            />
            <button className="w-full sm:w-auto px-6 py-2 text-white bg-green-800 hover:bg-green-900 rounded transition duration-300">
              Search
            </button>
          </form>
        </div>
      </div>

      <div className="">
        <section className="px-6 sm:px-10 lg:px-40 py-20 sm:py-36 mt-10 sm:mt-20 ">
          <div>
            <h2 className="text-3xl font-bold text-white font-crete">
              Products
            </h2>
            <p className="text-white  text-md font-medium font-mono py-3">
              Phones
            </p>
            <div className="grid lg:grid-cols-3 gap-5" id="servicecontainer">
              {products.map((product: any) => (
                <div
                  // data-aos="zoom-in"
                  key={product.asin}
                  className="transition  duration-700 ease-in-out hover:-translate-y-1 hover:scale-105 text-white py-20 bg-lime-950 border-transparent  border rounded-md shadow-md hover:shadow-xl px-6 sm:px-8 grid items-center"
                >
                  <Image  className="border w-full" src={product.productPhoto} width={24} height={24} alt="product photo" ></Image>
                  <h1 className="text-sm font-crete">{product.productTitle}</h1>
                  <p className="py-5 font-crete text-gray-300">
                    {product.productPrice}
                  </p>
                </div>
              ))}
              <div
                className={clsx({
                  "flex border justify-center bg-stone-800 shadow-md border-transparent rounded py-10":
                    !products,
                  hidden: products,
                })}
              >
                <div>
                  <p className="font-semibold text-xl">
                    There is no data to render
                  </p>
                </div>
              </div>
            </div>
          </div>
        </section>
      </div>
    </div>
    // flex border justify-center bg-stone-800 shadow-md border-transparent rounded py-10
  );
}

// text-white py-20 dark:bg-green-950 rounded-md border border-gray-600 shadow-md hover:shadow-xl mx-auto px-6 sm:px-8 grid items-center
