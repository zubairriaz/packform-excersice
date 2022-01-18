<template>
  <ag-grid-vue
    style="width: 100%; height: 300px"
    class="ag-theme-alpine"
    :columnDefs="columnDefs"
    :rowData="rowData"
  >
  </ag-grid-vue>
</template>

<script>
import "ag-grid-community/dist/styles/ag-grid.css";
import "ag-grid-community/dist/styles/ag-theme-alpine.css";
import { AgGridVue } from "ag-grid-vue3";
import { getColumns } from "../../../utils";

export default {
  name: "Grid",
  props: {
    items: Array,
  },
  watch: {
    items(newVal) {
      this.rowData = newVal;
    },
  },
  components: {
    AgGridVue,
  },
  setup(props) {
    return {
      columnDefs: getColumns(),
      rowData: props.items,
    };
  },
};
</script>
