"use client"

import clientApi from "@/api/clientApi"
import { Link } from "@/i18n/navigation"
import { useTranslations } from "next-intl";
import { FormEvent, useEffect, useRef, useState } from "react"

function debounce<T extends (...args: any[]) => void>(func: T, delay: number): T {
    let timer: NodeJS.Timeout;
    return ((...args: Parameters<T>) => {
        clearTimeout(timer);
        timer = setTimeout(() => func(...args), delay);
    }) as T;
}


export function Nav() {
    const [search, setSearch] = useState<any>(null)
    const [visible, setVisible] = useState<boolean>(false)

    const searchHandler = async (e: FormEvent<HTMLInputElement>) => {
        console.log(e)
        const input = e?.target as HTMLInputElement
        const data = await clientApi('/movies/search?q=' + encodeURIComponent(input?.value))
        setSearch(data?.data?.list)
        setVisible(true)
    }

    const searchRef = useRef<HTMLDivElement>(null)
    const inputRef = useRef<HTMLInputElement>(null)

    useEffect(() => {
        function handleClickOutside(event: MouseEvent) {

            if (inputRef.current && !inputRef.current.contains(event.target as Node) && (searchRef.current && !searchRef.current.contains(event.target as Node))) {
                setVisible(false);
            }
        }
        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };

    }, [])

    const t = useTranslations()

    return <nav className="bg-gradient-to-r from-purple-700 to-indigo-700 p-4 flex justify-between items-center shadow-lg text-white">
        <div className="text-3xl font-extrabold"><Link href="/">🎬 {t('Subtitle Hub')}</Link></div>
        <ul className="flex space-x-6 text-lg">
            <li><Link href="/" className="hover:text-gray-300">{t('Home')}</Link></li>
            <li><a href="/docs/index.html" className="hover:text-gray-300">{t('API')}</a></li>

        </ul>
        <div className="w-1/3">
            <input onClick={() => setVisible(!visible)} ref={inputRef} onInput={debounce(searchHandler, 800)} type="text" placeholder={t('Search') + "..."} className="w-full p-2 text-white rounded-lg border-2 border-white focus:outline-none focus:ring-2 focus:ring-white" />
            {visible && search?.length && <div className="flex"><div ref={searchRef} className={"overflow-y-auto h-[25rem] absolute bg-indigo-800 top-[4rem] p-4 z-10 text-white rounded-lg"}>
                <ul>
                    {search?.map((item: any) => <li className="flex flex-row gap-4 p-2 border-b-1 border-purple-200" key={item?.id}><img src={item?.poster} width={100} /><Link className="text-white" href={"/movie/" + item?.link_name}>{item?.name} {item?.year}</Link></li>)}
                </ul>
            </div></div>}
        </div>
    </nav>
}