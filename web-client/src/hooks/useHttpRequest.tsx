import { useState } from "react";

interface HttpRequestOptions {
  method: string;
  headers?: HeadersInit;
  body?: string;
}

interface HttpResponse<T> {
  data: T | null;
  loading: boolean;
  error?: Error | null;
  fetchData: any;
}

function useHttpRequest<T>(
  url: string,
  options: HttpRequestOptions
): HttpResponse<T> {
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<Error | null>(null);

  const fetchData = async () => {
    setLoading(true);

    try {
      const response = await fetch(url, options);
      const responseData = await response.json();

      setData(responseData);
      setLoading(false);
    } catch (err) {
      setError(err);
      setLoading(false);
    }
  };

  return { data, loading, fetchData, error };
}

export default useHttpRequest;
