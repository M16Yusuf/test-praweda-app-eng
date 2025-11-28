import { Input, Button, Table } from "antd";
import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router";

interface UserInterface {
  name: string;
  location: string;
  email: string;
  age: number;
  phone: string;
  cell: string;
  picture: string[];
}

function DataUserPage() {
  const [searchParams, setSearchParams] = useSearchParams({
    page: "1",
    results: "10",
  });
  const [search, setSearch] = useState("");
  const [data, setData] = useState<UserInterface[]>([]);
  const [loading, setLoading] = useState(false);

  // compute pagination values from URL params and data
  const pageSize = parseInt(searchParams.get("results") || "10", 10);
  const currentPage = parseInt(searchParams.get("page") || "1", 10);
  // Filter data client-side by name only (case-insensitive). This does not
  // trigger any backend fetch — it operates on the `data` already loaded.
  const filteredData = useMemo(() => {
    if (!search) return data;
    const q = search.trim().toLowerCase();
    return data.filter((d) => (d.name || "").toLowerCase().includes(q));
  }, [data, search]);

  // If backend returns exactly pageSize items, assume there may be a next page
  // and set total to allow AntD to enable the Next button. We base this on
  // the raw `data` (the current server page) so Next reflects server's page
  // availability. For display/current page we clamp against filteredData.
  let paginationTotal = (currentPage - 1) * pageSize + data.length;
  if (data.length === pageSize) paginationTotal = currentPage * pageSize + 1;

  // Do not clamp the current page to the filtered results. Keep `currentPage`
  // from the URL so pagination controls (Next/Prev) reflect server-driven
  // paging. The search filters only the currently loaded page (`data`).

  useEffect(() => {
    (async () => {
      setLoading(true);
      try {
        const res = await fetch(
          `http://127.0.0.1:8080/manipulasi?results=${searchParams.get(
            "results"
          )}&page=${searchParams.get("page")}`
        );
        if (!res.ok) throw new Error(`fetch error: ${res.status}`);
        const json = await res.json();

        // backend wraps payload in data, with nested data field
        const items = json?.data?.data ?? json?.data ?? json?.results ?? json;
        if (Array.isArray(items)) {
          setData(items as UserInterface[]);
        }
      } catch (err) {
        console.error("Failed to fetch manipulated data:", err);
      } finally {
        setLoading(false);
      }
    })();
  }, [searchParams]);

  return (
    <div
      style={{
        maxWidth: "1200px",
        justifySelf: "center",
        fontFamily: "sans-serif",
      }}
    >
      <h2 style={{ fontWeight: "800" }}>List</h2>
      <div
        style={{
          display: "flex",
          gap: "10px",
          marginBottom: "20px",
          justifyContent: "space-between",
        }}
      >
        <Input.Search
          placeholder="Search"
          style={{ maxWidth: "400px" }}
          variant="outlined"
          value={search}
          onChange={(e) => setSearch((e.target as HTMLInputElement).value)}
          disabled={loading}
        />
        <Button loading={loading}>+ New Data</Button>
      </div>
      <Table
        dataSource={filteredData}
        columns={[
          { title: "Nama", dataIndex: "name", key: "name" },
          { title: "Umur", dataIndex: "age", key: "age" },
          {
            title: "Alamat",
            dataIndex: "location",
            key: "location",
          },
          { title: "Email", dataIndex: "email", key: "email" },
          {
            title: "No telepon 1",
            dataIndex: "phone",
            key: "phone",
            minWidth: 120,
          },
          {
            title: "No telepon 2",
            dataIndex: "cell",
            key: "cell",
            minWidth: 120,
          },
          {
            title: "Gambar",
            dataIndex: "picture",
            key: "picture",
            render: (picture: string[]) => <img src={picture[1]} alt="User" />,
          },
        ]}
        rowKey="email"
        loading={loading}
        pagination={{
          pageSize,
          current: currentPage,
          total: paginationTotal,
          showSizeChanger: true,
          pageSizeOptions: ["5", "10", "20", "50"],
          onChange: (page, pageSize) => {
            setSearchParams({
              page: page.toString(),
              results: pageSize.toString(),
            });
          },
          // when user changes page size explicitly, go back to page 1
          onShowSizeChange: (_current, size) => {
            setSearchParams({ page: "1", results: size.toString() });
          },
          showTotal: () => `Items: ${paginationTotal}`,
        }}
      />
    </div>
  );
}

export default DataUserPage;
