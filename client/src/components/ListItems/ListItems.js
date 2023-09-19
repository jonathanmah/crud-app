import React, { Fragment, useEffect, useState } from "react";
import axios from "axios";
import DeleteRequest from "../DeleteItem/DeleteItem";
import EditItem from "../EditItem/EditItem";

const ListItems = () => {
  const [items, setItems] = useState([]);

  const getItems = async () => {
    try {
      const response = await axios.get("http://localhost:5000/items");
      setItems(response.data);
    } catch (error) {
      console.log("error");
    }
  };

  useEffect(() => {
    getItems();
  }, []);

  const DeleteItem = (item_id, e) => {
    setItems(items.filter((item) => item.item_id != item_id));
    DeleteRequest(item_id, e);
  };

  return (
    <Fragment>
      <table className="table mt-5 text-center">
        <thead>
          <tr>
            <th>Description</th>
            <th>Edit</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          {items.map((item) => (
            <tr key={item.item_id}>
              <td>{item.description}</td>
              <td>
                <EditItem item={item}></EditItem>
              </td>
              <td>
                <button
                  className="btn btn-danger"
                  onClick={(e) => DeleteItem(item.item_id, e)}
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </Fragment>
  );
};

export default ListItems;
