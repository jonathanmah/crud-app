import axios from "axios";

const DeleteRequest = async (item_id, e) => {
  e.preventDefault();
  try {
    const response = await axios.delete(
      `http://localhost:5000/items/${item_id}`
    );
  } catch (error) {
    console.log("error");
  }
};

export default DeleteRequest;
