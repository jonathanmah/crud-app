import axios from "axios";
import React, { useState, Fragment } from "react";
import { Modal } from "bootstrap";
import "bootstrap/dist/css/bootstrap.css";

const EditItem = ({ item }) => {
  const [description, setDescription] = useState(item.description);

  const UpdateItem = async (reqData, e) => {
    try {
      e.preventDefault();
      const response = await axios.put(
        `http://localhost:5000/items/${item.item_id}`,
        reqData
      );
      window.location = "/";
    } catch (error) {
      console.log("error");
    }
  };

  return (
    <Fragment>
      <button
        type="button"
        className="btn btn-primary"
        data-bs-toggle="modal"
        data-bs-target={`#id${item.item_id}`}
      >
        Edit
      </button>
      <div
        className="modal fade"
        id={`id${item.item_id}`}
        tabIndex="-1"
        aria-hidden="true"
      >
        <div className="modal-dialog">
          <div className="modal-content">
            <div className="modal-header">
              <h1 className="modal-title fs-5">Edit Item</h1>
              <button
                type="button"
                className="btn-close"
                data-bs-dismiss="modal"
                aria-label="Close"
                onClick={() => setDescription(item.description)}
              ></button>
            </div>
            <div className="modal-body">
              <input
                type="text"
                className="form-control"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
              ></input>
            </div>
            <div className="modal-footer">
              <button
                type="button"
                className="btn btn-warning"
                data-bs-dismiss="modal"
                onClick={(e) => {
                  UpdateItem({ description: description }, e);
                }}
              >
                Save
              </button>
              <button
                type="button"
                className="btn btn-danger"
                data-bs-dismiss="modal"
                onClick={() => setDescription(item.description)}
              >
                Close
              </button>
            </div>
          </div>
        </div>
      </div>
    </Fragment>
  );
};

export default EditItem;
