"use client"
import { Link } from "@/i18n/navigation";
import { useTranslations } from "next-intl";
import { useSearchParams } from "next/navigation";

export default function Pagination({ pagination, module }: any) {
    const t = useTranslations();

    const params = new URLSearchParams(useSearchParams().toString());
    params.delete("page")
    const qs = params.toString() ? "&" + params.toString() : ""

    return <div className="px-4 bottom-0 right-0 items-center w-full py-4  sm:flex sm:justify-between">
        <div className="flex items-center mb-4 sm:mb-0">
            <Link href={module + "?page=" + pagination?.prev + qs} className="inline-flex justify-center p-1 text-indigo-500 rounded cursor-pointer hover:text-indigo-900 hover:bg-indigo-100 dark:text-indigo-400 dark:hover:bg-indigo-700 dark:hover:text-white">
                <svg className="rtl:hidden w-7 h-7" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clipRule="evenodd" /></svg>

                <svg className="ltr:hidden w-7 h-7" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clipRule="evenodd" /></svg>
            </Link>
            <Link href={module + "?page=" + pagination?.next + qs} className="inline-flex justify-center p-1 ltr:mr-2 ms-2 text-indigo-500 rounded cursor-pointer hover:text-indigo-900 hover:bg-indigo-100 dark:text-indigo-400 dark:hover:bg-indigo-700 dark:hover:text-white">
                <svg className="rtl:hidden w-7 h-7" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clipRule="evenodd" /></svg>

                <svg className="ltr:hidden w-7 h-7" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clipRule="evenodd" /></svg>

            </Link>
            <span className="text-sm font-normal text-indigo-500 dark:text-indigo-400">{t("Showing")} <span className="font-semibold text-indigo-900 dark:text-white">{pagination?.from}-{pagination?.to}</span> {t("of")} <span className="font-semibold text-indigo-900 dark:text-white">{pagination?.count}</span></span>
        </div>
        <div className="flex items-center">
            <Link href={module + "?page=" + pagination?.prev + qs} className="ms-2 inline-flex items-center justify-center flex-1 px-3 py-2 text-sm font-medium text-center border border-indigo-300 text-indigo-700 dark:text-indigo-200 rounded-lg bg-indigo-100 focus:ring-4 focus:ring-primary-300 dark:bg-indigo-600 dark:hover:bg-indigo-600 dark:focus:ring-indigo-800">
                <svg className="rtl:hidden w-5 h-5 ltr:mr-1 ms-1 -ml-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clipRule="evenodd" /></svg>

                <svg className="ltr:hidden w-5 h-5 ltr:ml-1 rtl:mr-1 -mr-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clipRule="evenodd" /></svg>

                {t("Previous")}
            </Link>
            <Link href={module + "?page=" + pagination?.next + qs} className="ms-2 inline-flex items-center justify-center flex-1 px-3 py-2 text-sm font-medium text-center border border-indigo-300 text-indigo-700 dark:text-indigo-200 rounded-lg bg-indigo-100 focus:ring-4 focus:ring-primary-300 dark:bg-indigo-600 dark:hover:bg-indigo-700 dark:focus:ring-indigo-800">
                {t("Next")}
                <svg className="rtl:hidden w-5 h-5 ltr:ml-1 rtl:mr-1 -mr-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clipRule="evenodd" /></svg>

                <svg className="ltr:hidden w-5 h-5 ltr:mr-1 ms-1 -ml-1" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clipRule="evenodd" /></svg>
            </Link>
        </div>
    </div>
}