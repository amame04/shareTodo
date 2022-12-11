flatpickr("#flatpickr", {
  locale: 'ja',
  enableTime: true,
  dateFormat: "Y-m-d H:i",
});

$(function() {
  $("#select").multipleSelect();
});