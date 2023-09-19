import React, { Fragment, useState } from "react";
import axios from "axios";
import "./InputItem.css";

const InputItem = () => {
  const [description, setDescription] = useState("");

  const onSubmitForm = async (e) => {
    e.preventDefault();
    try {
      const res = await axios.post("http://localhost:5000/items", {
        description: description,
      });
      setDescription("");
      window.location = "/";
    } catch (error) {
      console.error(error.message);
    }
  };

  return (
    <Fragment>
      <form className="d-flex input-bar">
        <input
          type="text"
          className="form-control"
          placeholder="Enter an item"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <button className="btn btn-success" onClick={(e) => onSubmitForm(e)}>
          Add
        </button>
      </form>
    </Fragment>
  );
};

export default InputItem;
