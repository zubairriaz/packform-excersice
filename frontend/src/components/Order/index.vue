<template>
  <div id="Order">
    <div>
      <span class="label">Search</span>

      <input class="search-input" v-model="searchText" />
    </div>
    <div class="date-picker">
      <span class="label">Created Date</span>
      <Datepicker
        v-model="date"
        range
        twoCalendars
        twoCalendarsSolo
        placeholder="Start Date - End Date"
        @update:modelValue="onDateSelected"
      />
    </div>

    <grid :items="orders" v-if="orders.length > 0"></grid>
    <span class="v-paginate"
      ><pagination
        v-model="page"
        :records="totalRecords"
        :per-page="recordPerPage"
        @paginate="onPageChange"
      />
    </span>
  </div>
</template>

<script>
import Grid from "./Grid/grid.vue";
import { AXIOS } from "../../utils";
import Pagination from "v-pagination-3";
import { RECORDS_PER_PAGE } from "../../utils";
import Datepicker from "vue3-date-time-picker";
import "vue3-date-time-picker/dist/main.css";

import { ref } from "vue";

export default {
  name: "Order",
  setup() {
    const date = ref();

    return {
      date,
    };
  },
  data() {
    return {
      orders: [],
      page: 1,
      recordPerPage: RECORDS_PER_PAGE,
      searchText: "",
    };
  },
  watch: {
    searchText(newVal) {
      if (this.timeOut) {
        clearTimeout(this.timeOut);
      }
      this.timeOut = setTimeout(
        () => this.getOrders({ start: 0, searchText: newVal }),
        500
      );
    },
  },
  computed: {
    totalRecords() {
      return this.orders.length > 0 ? this.orders[0].FullCount : 0;
    },
  },
  methods: {
    getOrders: function ({
      start,
      startDate = "",
      endDate = "",
      searchText = "",
    }) {
      if (!startDate) {
        startDate = new Date("1/1/0001 12:00:00").toISOString();
      }
      if (!endDate) {
        endDate = new Date("1/1/9999 12:00:00").toISOString();
      }
      AXIOS.get(
        `http://localhost:8001/api/orders?count=${this.recordPerPage}&start=${start}&startDate=${startDate}&endDate=${endDate}&searchText=${searchText}`
      )
        .then((response) => {
          this.orders = [...response.data];
        })
        .catch(function (error) {
          console.log(error);
        });
    },
    onPageChange(pageNumber) {
      console.log(pageNumber, this.page);
      if (pageNumber == 1) {
        let start = 0;
        this.getOrders({ start });
      } else {
        let start = pageNumber * 5 - 5;
        this.getOrders({ start });
      }
    },
    onDateSelected(date) {
      let myTarget = JSON.parse(JSON.stringify(date));
      const startDate = myTarget[0];
      const endDate = myTarget[1];
      this.getOrders({ start: 1, startDate, endDate, searchText: "" });
    },
  },
  components: {
    Grid,
    Pagination,
    Datepicker,
  },
  beforeMount() {
    this.getOrders({ start: 0 });
  },
};
</script>
<style>
ul {
  display: flex;
  flex-direction: row;
  list-style-type: none;
  list-style: none;
}
.date-picker div{
  margin-top: 4px;
}
.date-picker {
  margin: 20px 0;
  display: flex;
  flex-direction: row;
}
.label {
  margin: 20px;
  font-weight: 500;
}

.search-input{
  height: 25px;
}
</style>
