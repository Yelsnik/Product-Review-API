export default function Page() {
  
  return (
    <div className="mt-20 py-5">
      <div className="px-20">
        <h1>Top Ten Products</h1>
      </div>
      <div className="">
        <nav className=" border border-gray-800 border-transparent fixed top-0 left-0 right-0 z-50 bg-opacity-75 backdrop-filter backdrop-blur-lg items-center px-6 sm:px-10 lg:px-20 py-5 text-white bg-transparent">
          <div>Leaderboard</div>
        </nav>
      </div>

      <section className="mt-10">
        <div className="grid mt-20 px-20">
          {/* <ol className="list-decimal">
            {products.map((product: any) => (
              <li
                key={product.asin}
                className="px-2 transition border-stone-900 duration-700 ease-in-out hover:-translate-y-1 hover:scale-100 border bg-stone-700 shadow-md hover:shadow-xl py-5 rounded grid mt-5"
              >
                <div className="grid">
                  <h1>{product.productTitle}</h1>
                  <div>
                    <button className="border px-2 py-2 rounded bg-green-900 border-transparent transition duration-700 ease-in-out hover:-translate-y-1 hover:scale-105  mt-5 text-start">
                      view details
                    </button>
                  </div>
                </div>
              </li>
            ))}
          </ol> */}
        </div>
      </section>
    </div>
  );
}
