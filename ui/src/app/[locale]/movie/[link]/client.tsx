"use client"
import clientApi from "@/api/clientApi";
import { useTranslations } from "next-intl";
import { useEffect, useRef, useState } from "react";

export function ClientMovie({ movie }: any) {

    const langs = [...new Set(movie?.subtitles.map((item: any) => item.lang))];
    const [subtitles, setSubtitles] = useState<any>(movie?.subtitles)
    const t = useTranslations()
    const langRef = useRef<HTMLSelectElement>(null)
    useEffect(() => {
        const value = localStorage.getItem('lang') || ""
        if (langRef.current && value != "") {
            const options = Array.from(langRef.current.options).map((option) => option.value)

            if (options.includes(value)) {

                langRef.current.value = value

                const event = new Event('change', { 'bubbles': true });
                langRef?.current.dispatchEvent(event);
            }
        }
    }, [])
    return <>
        <div className="container mx-auto p-6" >
            <div className="grid grid-cols-1 md:grid-cols-5 gap-6">
                <div>
                    <img src={movie?.poster} alt={movie?.title} className="w-full rounded-lg shadow-lg" />
                </div>
                <div className="md:col-span-2">
                    <h2 className="text-4xl font-extrabold mb-4 text-gray-500">{movie?.name} {movie?.year}</h2>
                    <p className="mb-2 text-gray-500 text-lg">{t("Year")}: {movie?.year}</p>
                    <div className="flex items-center space-x-3 mt-2">
                        <a href={"https://www.imdb.com/title/" + movie?.imdb_code} className="text-blue-400 text-lg">🔗 IMDB</a>
                    </div>
                </div>
            </div>


            <h3 className="text-3xl font-bold mb-3 text-indigo-400 mt-6">{t('Download Subtitles')}</h3>
            <select ref={langRef} onChange={(e) => {
                if (e.currentTarget.value != "all") {
                    setSubtitles(movie?.subtitles?.filter((item: any) => item?.lang == e.currentTarget.value))
                } else {
                    setSubtitles(movie?.subtitles)
                }

                if (e.currentTarget.value) {
                    localStorage.setItem('lang', e.currentTarget.value)
                }

            }} className="p-3 mb-4 text-black rounded-lg bg-gray-200 focus:outline-none focus:ring-2 focus:ring-indigo-500">
                <option value="all">{t('All Languages')}</option>
                {langs?.map((item: any) => <option key={item} value={item}>{item}</option>)}
            </select> {typeof window !== 'undefined' && localStorage.getItem('lang')}
            <table className="space-y-3 w-full">
                <tbody>
                    {subtitles?.map((item: any) => (
                        <tr key={item?.id} className="bg-gray-800 p-4 rounded-lg shadow-md justify-between items-center hover:bg-gray-700 transition-colors border-b-1 border-white">
                            <td className="text-white text-lg p-4">{item?.title}</td>
                            <td className="p-4 text-white">{item?.lang}</td>
                            <td className="p-4" width={'10%'}><a onClick={async (e) => {
                                if (item?.file_name == "") {
                                    e.preventDefault()
                                    const api = await clientApi('/subtitles/' + item?.id + '/download')
                                    if (api.status == 1) {
                                        (e.target as HTMLLinkElement).href = api?.data?.subtitle?.file_name

                                        const event = new Event('click', { 'bubbles': true });
                                        e.target.dispatchEvent(event);

                                    }

                                }
                            }} href={process.env.NEXT_PUBLIC_BASE_URL + "files/subtitles/" + item?.file_name} className={(item?.file_name == "" ? "!text-red-500 " : "") + "text-indigo-400 hover:text-indigo-300"}>{t('Download')}</a>
                            </td>
                        </tr>)
                    )}
                </tbody>
            </table>
        </div ></>
}