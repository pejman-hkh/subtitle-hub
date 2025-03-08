import Pagination from "@/components/pagination";
import { Link } from "@/i18n/navigation";
import { useTranslations } from "next-intl";

export default function ClientIndex({ data }: any) {
    const movies = data?.data?.list
    const t = useTranslations()

    return <div className="container mx-auto p-6">

        <h2 className="text-4xl font-extrabold mb-6 text-indigo-400">{t('Latest Movies')}</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 gap-6">
            {movies?.map((item: any) => <Link href={"/movie/" + item?.link_name} key={item?.id} className="bg-gray-800 p-4 rounded-lg shadow-lg hover:scale-105 transition-transform">
                <img src={item?.poster} alt={item?.name} className="w-full rounded-md" />
                <h3 className="text-xl font-semibold mt-3 text-white">{item?.name} {item?.year}</h3>
                <div className="flex items-center space-x-3 mt-2">
                    <a target="_blank" href={"https://www.imdb.com/title/" + item?.imdb_code + "/"} className="text-blue-400 text-lg">🔗 IMDB</a>
                </div>

            </Link>)}

        </div>

        <Pagination pagination={data?.data?.pagination} module={"/"} />
    </div>
}