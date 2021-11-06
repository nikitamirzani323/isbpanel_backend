<script>
    export let card_title = "";
    export let card_footer = 0;
   
    let height = 0;

    $:{
        height = window.innerHeight - 250;
    }
</script>
<div class="card mt-1">
    <div class="card-header">
        <h5 class="card-title">{card_title}</h5>
    </div>
    <div class="card-body overflow-auto" style="padding: 5px;margin:0px;height:{height}px;">
        <div class="table-responsive">
            <slot name="card-body" />
        </div>
    </div>
    <div class="card-footer">Record : {card_footer}</div>
</div>