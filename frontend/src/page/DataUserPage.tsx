import { Input, Button, Table } from "antd";
import { useState } from "react";

interface DummyUser {
  name: string;
  location: string;
  email: string;
  age: number;
  phone: string;
  cell: string;
  picture: string[];
}

const dummyData: DummyUser[] = [
  {
    name: "Miss Polyana Kotenko",
    location:
      "3967 Holmogorskiy provulok, Bogoduhiv, Kirovogradska, Ukraine, 20241",
    email: "polyana.kotenko@example.com",
    age: 26,
    phone: "(097) M41-1541",
    cell: "(097) F05-1370",
    picture: [
      "https://randomuser.me/api/portraits/women/76.jpg",
      "https://randomuser.me/api/portraits/med/women/76.jpg",
      "https://randomuser.me/api/portraits/thumb/women/76.jpg",
    ],
  },
  {
    name: "Mrs Neea Peura",
    location: "4206 Hämeenkatu, Pori, Kymenlaakso, Finland, 76462",
    email: "neea.peura@example.com",
    age: 77,
    phone: "08-404-835",
    cell: "049-672-39-12",
    picture: [
      "https://randomuser.me/api/portraits/women/66.jpg",
      "https://randomuser.me/api/portraits/med/women/66.jpg",
      "https://randomuser.me/api/portraits/thumb/women/66.jpg",
    ],
  },
  {
    name: "Ms Amalie Olsen",
    location: "5497 Stormgade, Juelsminde, Nordjylland, Denmark, 59955",
    email: "amalie.olsen@example.com",
    age: 65,
    phone: "07380543",
    cell: "28365825",
    picture: [
      "https://randomuser.me/api/portraits/women/89.jpg",
      "https://randomuser.me/api/portraits/med/women/89.jpg",
      "https://randomuser.me/api/portraits/thumb/women/89.jpg",
    ],
  },
  {
    name: "Miss Mina Glavaš",
    location: "8056 Ante Lambaše, Pirot, Bor, Serbia, 65703",
    email: "mina.glavas@example.com",
    age: 54,
    phone: "011-6612-178",
    cell: "067-0613-203",
    picture: [
      "https://randomuser.me/api/portraits/women/59.jpg",
      "https://randomuser.me/api/portraits/med/women/59.jpg",
      "https://randomuser.me/api/portraits/thumb/women/59.jpg",
    ],
  },
  {
    name: "Mr Sherif Van Lee",
    location:
      "3595 Achter Het Zwarte Beerke, Zevenhuizen Zh, Groningen, Netherlands, 9982 PX",
    email: "sherif.vanlee@example.com",
    age: 69,
    phone: "(0281) 294686",
    cell: "(06) 22761217",
    picture: [
      "https://randomuser.me/api/portraits/men/86.jpg",
      "https://randomuser.me/api/portraits/med/men/86.jpg",
      "https://randomuser.me/api/portraits/thumb/men/86.jpg",
    ],
  },
  {
    name: "Mrs Florence Taylor",
    location: "5576 Concession Road 23, Lumsden, Yukon, Canada, R5Q 6Z1",
    email: "florence.taylor@example.com",
    age: 25,
    phone: "C38 U64-4159",
    cell: "P12 W97-7625",
    picture: [
      "https://randomuser.me/api/portraits/women/13.jpg",
      "https://randomuser.me/api/portraits/med/women/13.jpg",
      "https://randomuser.me/api/portraits/thumb/women/13.jpg",
    ],
  },
  {
    name: "Ms Draginja Šarić",
    location: "1049 Lukićeva, Zaječar, Raška, Serbia, 63466",
    email: "draginja.saric@example.com",
    age: 52,
    phone: "030-3117-234",
    cell: "064-6519-459",
    picture: [
      "https://randomuser.me/api/portraits/women/39.jpg",
      "https://randomuser.me/api/portraits/med/women/39.jpg",
      "https://randomuser.me/api/portraits/thumb/women/39.jpg",
    ],
  },
  {
    name: "Mr سهیل محمدخان",
    location: "4392 پیروزی, اهواز, گیلان, Iran, 29363",
    email: "shyl.mhmdkhn@example.com",
    age: 60,
    phone: "084-33632013",
    cell: "0987-307-7644",
    picture: [
      "https://randomuser.me/api/portraits/men/36.jpg",
      "https://randomuser.me/api/portraits/med/men/36.jpg",
      "https://randomuser.me/api/portraits/thumb/men/36.jpg",
    ],
  },
  {
    name: "Mr Daniel Christiansen",
    location: "7719 Stormgade, Ugerløse, Midtjylland, Denmark, 31085",
    email: "daniel.christiansen@example.com",
    age: 66,
    phone: "85212556",
    cell: "81283202",
    picture: [
      "https://randomuser.me/api/portraits/men/6.jpg",
      "https://randomuser.me/api/portraits/med/men/6.jpg",
      "https://randomuser.me/api/portraits/thumb/men/6.jpg",
    ],
  },
  {
    name: "Miss Vladislava da Cunha",
    location: "6731 Rua Boa Vista , Pindamonhangaba, Tocantins, Brazil, 72611",
    email: "vladislava.dacunha@example.com",
    age: 38,
    phone: "(89) 9376-0454",
    cell: "(16) 1886-4055",
    picture: [
      "https://randomuser.me/api/portraits/women/92.jpg",
      "https://randomuser.me/api/portraits/med/women/92.jpg",
      "https://randomuser.me/api/portraits/thumb/women/92.jpg",
    ],
  },
];

function DataUserPage() {
  const [data] = useState(dummyData);
  return (
    <div style={{ maxWidth: "1200px", justifySelf: "center" }}>
      <h1>List</h1>
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
          style={{ maxWidth: "600px" }}
          variant="filled"
        />
        <Button>+ New Data</Button>
      </div>
      <Table
        dataSource={data}
        columns={[
          { title: "Nama", dataIndex: "name", key: "name" },
          { title: "Umur", dataIndex: "age", key: "age" },
          { title: "Alamat", dataIndex: "location", key: "location" },
          { title: "Email", dataIndex: "email", key: "email" },
          { title: "No telepon 1", dataIndex: "phone", key: "phone" },
          { title: "No telepon 2", dataIndex: "cell", key: "cell" },
          {
            title: "Gambar",
            dataIndex: "picture",
            key: "picture",
            render: (picture: string[]) => <img src={picture[2]} alt="User" />,
          },
        ]}
        rowKey="email"
      />
    </div>
  );
}

export default DataUserPage;
