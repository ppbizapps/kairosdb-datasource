export class KairosDBConfigCtrl {
  static templateUrl = 'partials/config.html';
  datasourceSrv: any;
  current: any;

  /** @ngInject */
  constructor($scope, datasourceSrv) {
    this.datasourceSrv = datasourceSrv;
    this.current.jsonData = this.current.jsonData || {};
    if (Object.keys(this.current.jsonData).length === 0) {
      this.current.jsonData.selectedDataSources = []
      this.current.jsonData.multi = false
    }

    this.getAllDataSources()
    this.getAllKairosDataSources()
  }

  getAllDataSources() {
    this.current.jsonData.allDataSources = this.datasourceSrv.getAll()
  }

  getAllKairosDataSources() {
    this.current.jsonData.allKairosDataSources = []
    for (let key of Object.keys(this.current.jsonData.allDataSources)) {
      const ds = this.current.jsonData.allDataSources[key]
      if (ds.type == 'grafana-kairosdb-datasource' && !ds.jsonData.multi) {
        this.current.jsonData.allKairosDataSources.push(ds)
      }
    }
  }

  selectDataSource(ds) {
    if (ds) this.current.jsonData.selectedDataSources.push(ds)
  }



}