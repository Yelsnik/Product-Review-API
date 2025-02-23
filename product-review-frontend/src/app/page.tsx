// import Image from "next/image";

import Link from "next/link";

export default function Home() {
  return (
    <div>
      <nav className="border border-gray-800 border-transparent fixed top-0 left-0 right-0 z-50 bg-opacity-75 backdrop-filter backdrop-blur-lg items-center px-6 sm:px-10 lg:px-40 py-5 text-white bg-transparent">
        <div className="flex justify-between">
          <div>Product Review</div>
          <Link href={"leaderboard"}>Leaderboard</Link>
        </div>
      </nav>

      <div className="mt-40 px-40">
        <div className="px-40">
          <form action="" className="flex justify-center w-full gap-5">
            <div className="">
              <input className="border-transparent w-72 py-2 px-7 font-semibold items-center text-black outline-none bg-gray-100 rounded" type="text" placeholder="search for a product" />
            </div>

            <div className=""><button className="border hover:bg-green-900 rounded px-5 py-2 bg-green-800 border-transparent ">Search</button></div>
          </form>
        </div>
      </div>

      <div className="">
        <section className="px-6 sm:px-10 lg:px-40 py-20 sm:py-36 mt-10 sm:mt-20 ">
          <div>
            <h2 className="text-3xl font-bold text-white font-crete">
              Products
            </h2>
            <p className="text-white text-md font-medium font-mono py-3">
              No products
            </p>
            <div
              className="grid  gap-5"
              id="servicecontainer"
            >
              {/* <div
                // data-aos="zoom-in"
                className="text-white py-20 border rounded-md shadow-md hover:shadow-xl px-6 sm:px-8 grid items-center"
              >
                <h1 className="text-2xl font-crete">Iphone</h1>
                <p className="py-5 font-crete text-gray-300">A good phone</p>
              </div>
              <div
                // data-aos="zoom-in"
                className="text-white py-20 border rounded-md shadow-md hover:shadow-xl px-6 sm:px-8 grid items-center"
              >
                <h1 className="text-2xl font-crete">Iphone</h1>
                <p className="py-5 font-crete text-gray-300">A good phone</p>
              </div>
              <div
                // data-aos="zoom-in"
                className="text-white py-20 border rounded-md shadow-md hover:shadow-xl px-6 sm:px-8 grid items-center"
              >
                <h1 className="text-2xl font-crete">Iphone</h1>
                <p className="py-5 font-crete text-gray-300">A good phone</p>
              </div> */}
              <div className="flex border justify-center bg-stone-800 shadow-md border-transparent rounded py-10">
                <div>
                  <p className="font-semibold text-xl">There is no data to render</p>
                </div>
              </div>
            </div>
          </div>
        </section>
      </div>
    </div>
  );
}

// text-white py-20 dark:bg-green-950 rounded-md border border-gray-600 shadow-md hover:shadow-xl mx-auto px-6 sm:px-8 grid items-center