import { useState, useContext } from "react";
import { AuthContext } from "../Context/AuthContext";

export default function Signup() {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
    email: "",
  });

  const { signup } = useContext(AuthContext);

  const handleSubmit = async (e) => {
    e.preventDefault();
    signup(formData)
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="mb-3">
        <label htmlFor="Email1" className="form-label">
          Email address
        </label>
        <input
          type="email"
          className="form-control"
          id="Email1"
          aria-describedby="emailHelp"
          value={formData.email}
          onChange={handleChange}
          name="email"
        />
        <div id="emailHelp" className="form-text">
          We'll never share your email with anyone else.
        </div>
      </div>
      <div className="mb-3">
        <label htmlFor="username" className="form-label">
          Username
        </label>
        <input
          type="text"
          className="form-control"
          id="username"
          aria-describedby="UsernameHelp"
          value={formData.username}
          onChange={handleChange}
          name="username"
        />
      </div>
      <div className="mb-3">
        <label htmlFor="exampleInputPassword1" className="form-label">
          Password
        </label>
        <input
          type="password"
          className="form-control"
          id="exampleInputPassword1"
          value={formData.password}
          onChange={handleChange}
          name="password"
        />
      </div>
      <button type="submit" className="btn btn-primary">
        Submit
      </button>
    </form>
  );
}
