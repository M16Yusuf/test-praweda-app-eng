const warna: string[] = ["merah", "kuning", "hijau", "pink", "ungu"];
const pakaian: string[] = ["baju", "celana", "topi", "jaket", "sepatu"];
const statusDiskon: string[] = ["Diskon", "Sale", "Diskon", "Sale", "Sale"];

const manipulasiArray: string[] = [];

warna.push("maroon");

for (let i = 0; i < warna.length; i++) {
  const pakaianItem = pakaian[i] || pakaian[pakaian.length - i];
  const diskonItem = statusDiskon[i] || statusDiskon[statusDiskon.length - i];

  manipulasiArray.push(`${warna[i]} ${pakaianItem} ${diskonItem}`);
}

console.log(manipulasiArray);

export default manipulasiArray;
