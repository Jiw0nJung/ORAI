// 모듈추가
const express = require("express");
const app = express();
var bodyParser = require("body-parser");
// 하이퍼레저 모듈추가+연결속성파일로드
const { FileSystemWallet, Gateway } = require("fabric-network");
const fs = require("fs");
const path = require("path");
const ccpPath = path.resolve(__dirname, "..", "network", "connection.json");
const ccpJSON = fs.readFileSync(ccpPath, "utf8");
const ccp = JSON.parse(ccpJSON);
// 서버속성
const PORT = 8080;
const HOST = "0.0.0.0";
// app.use
app.use(express.static(path.join(__dirname, "views")));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
// 라우팅
// 1. 페이지라우팅
app.get("/", (req, res) => {
  res.sendFile(__dirname + "/index.html");
});
app.get("/adduser", (req, res) => {
  res.sendFile(__dirname + "/adduser.html");
});
app.get("/invest", (req, res) => {
  res.sendFile(__dirname + "/invest.html");
});
app.get("/query", (req, res) => {
  res.sendFile(__dirname + "/query.html");
});
// 2. REST라우팅

// 2.1 adduser
app.post("/user", async (req, res) => {
  const name = req.body.name;

  const walletPath = path.join(process.cwd(), "wallet");
  const wallet = new FileSystemWallet(walletPath);
  console.log(`Wallet path: ${walletPath}`);

  const userExists = await wallet.exists("user1");
  if (!userExists) {
    console.log(
      'An identity for the user "user1" does not exist in the wallet'
    );
    console.log("Run the registerUser.js application before retrying");
    return;
  }
  const gateway = new Gateway();
  await gateway.connect(ccp, {
    wallet,
    identity: "user1",
    discovery: { enabled: false }
  });
  const network = await gateway.getNetwork("mychannel");
  const contract = network.getContract("reitscc");
  await contract.submitTransaction("addUser", name);
  console.log("Transaction has been submitted");
  await gateway.disconnect();

  res.status(200).send("Transaction has been submitted");
});
// 2.2 invest
app.post("/move", async (req, res) => {
  const username = req.body.username;
  const projectname = req.body.projectname;
  const value = req.body.value;

  const walletPath = path.join(process.cwd(), "wallet");
  const wallet = new FileSystemWallet(walletPath);
  console.log(`Wallet path: ${walletPath}`);

  const userExists = await wallet.exists("user1");
  if (!userExists) {
    console.log(
      'An identity for the user "user1" does not exist in the wallet'
    );
    console.log("Run the registerUser.js application before retrying");
    return;
  }
  const gateway = new Gateway();
  await gateway.connect(ccp, {
    wallet,
    identity: "user1",
    discovery: { enabled: false }
  });
  const network = await gateway.getNetwork("mychannel");
  const contract = network.getContract("reitscc");
  await contract.submitTransaction("invest", username, projectname, value);
  console.log("Transaction has been submitted");
  await gateway.disconnect();

  res.status(200).send("Transaction has been submitted");
});

// 2.3 query
app.get("/user", async (req, res) => {
  const name = req.query.name;

  const walletPath = path.join(process.cwd(), "wallet");
  const wallet = new FileSystemWallet(walletPath);
  console.log(`Wallet path: ${walletPath}`);
  const userExists = await wallet.exists("user1");
  if (!userExists) {
    console.log(
      'An identity for the user "user1" does not exist in the wallet'
    );
    console.log("Run the registerUser.js application before retrying");
    return;
  }
  const gateway = new Gateway();
  await gateway.connect(ccp, {
    wallet,
    identity: "user1",
    discovery: { enabled: false }
  });
  const network = await gateway.getNetwork("mychannel");
  const contract = network.getContract("reitscc");
  const result = await contract.evaluateTransaction("query", name);
  console.log(
    `Transaction has been evaluated, result is: ${result.toString()}`
  );

  var obj = JSON.parse(result);
  res.status(200).json(obj);
});

// 서버시작
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
