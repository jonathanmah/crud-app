import ListItems from "../components/ListItems/ListItems";
import Header from "../components/Header/Header";
import InputItem from "../components/InputItem/InputItem";
import "bootstrap/dist/css/bootstrap.css";
import Gopher from "../img/gopher.jpg";

const App = () => {
  let headerText = "Go CRUD App";
  return (
    <div>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <div
          style={{ display: "flex", alignItems: "center", marginLeft: "-90px" }}
        >
          <img
            src={Gopher}
            alt="Gopher"
            style={{ width: "200px", height: "auto", marginRight: "-50px" }}
          />
          <Header text={headerText} />
        </div>
      </div>

      <br />
      <div className="container">
        <InputItem></InputItem>
        <br />
        <ListItems></ListItems>
      </div>
    </div>
  );
};

export default App;
