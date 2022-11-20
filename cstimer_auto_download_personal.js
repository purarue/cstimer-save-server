// ==UserScript==
// @name           CSTimer Auto Download
// @namespace      https://greasyfork.org/en/users/96096
// @match          https://cstimer.net/*
// @match          http://localhost:4633/*
// @grant          none
// @version        0.2
// @author         Sean Breckenridge
// @description    Companion script for cstimer-save-server to update a local file with my cstimer.net solves
// @license        GPL3
// ==/UserScript==

(() => {
  // redirect to local server
  if (window.location.href.includes("cstimer.net")) {
    window.location.href = "http://localhost:4633"
  }
  // configuration
  const PORT = 8553;
  const SECRET = "";

  async function runExport() {
    return storage.exportAll().then(function (exportObj) {
      exportObj["properties"] = mathlib.str2obj(localStorage["properties"]);
      return JSON.stringify(exportObj);
    });
  }

  runExport().then((data) =>
    fetch(`http://127.0.0.1:${PORT}/?auth=${SECRET}`, {
      method: "post",
      body: data,
    })
  );
})();

