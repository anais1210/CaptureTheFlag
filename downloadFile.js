const ScanForm = () => {
  const [address, setAddress] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault();

    const goURL = "http://34.77.36.161:3942/?secretKey=" + secretKey;
    graphlresultaddr.innerHTML = "Getting Data...";

    fetch(goURL, {
      method: "POST",
      headers: {
        "Content-Type": "multipart/form-data",
      },
      body: "secretKey=" + secretKet,
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        graphlresulttxcount.innerHTML = data;
      });

    // body: 'query { contracts (where: {address: {_eq: "' + address + '"}}) { address transactions } operations (where: {from: {_eq: "' + address + '"}, _or: {to: {_eq: "' + address + '"}}}) {from hash to value timestamp}}'
  };
};
