import { useState } from "react";

const Search = () => {
  const [searchText, setSearchText] = useState("");

  const updateSearchText = (e) => {
    setSearchText(e.target.value);
  };
  function handleOnClick(e) {
    e.PreventDefault();
  }
  return (
    <div className="input-group">
      <input
        type="text"
        className="form-control"
        placeholder="URL of the website that needs to be verified."
        value={searchText}
        onChange={updateSearchText}
      />
      <div className="input-group-append">
        <button
          className="btn btn-outline-secondary"
          type="button"
          onClick={handleOnClick}
        >
          Check
        </button>
      </div>
    </div>
  );
};

export default Search;
