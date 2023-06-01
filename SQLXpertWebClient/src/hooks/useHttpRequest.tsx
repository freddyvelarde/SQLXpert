import { useState } from "react";

interface HttpRequestOptions {
  method: string;
  headers?: HeadersInit;
  body?: string;
}

interface HttpResponse<T> {
  data: T | null;
  loading: boolean;
  error?: boolean | null;
  fetchData: any;
}

function useHttpRequest<T>(
  url: string,
  options: HttpRequestOptions
): HttpResponse<T> {
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<boolean | null>(null);

  const fetchData = async () => {
    setLoading(true);

    try {
      const response = await fetch(url, options);
      const responseData = await response.json();

      setData(responseData);
      setLoading(false);
      setError(false);
    } catch (err) {
      setError(true);
      setLoading(false);
    }
  };

  return { data, loading, fetchData, error };
}

export default useHttpRequest;
