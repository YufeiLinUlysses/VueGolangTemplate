<template>
    <div class="btn-group-vertical">
        <router-link :to="{name: 'Financial'}">
            <button type="button" class="btn btn-secondary custom">Financial</button>
        </router-link>
        <router-link :to="{name: 'HealthInfo'}">
            <button type="button" class="btn btn-secondary custom">Health Information</button>
        </router-link>
        <router-link :to="{name: 'QualityMeasure'}">
            <button type="button" class="btn btn-secondary custom">Quality Measure</button>
        </router-link>
        <router-link :to="{name: 'SocialDeter'}">
            <button type="button" class="btn btn-secondary custom">Social Determinant</button>
        </router-link>
        <router-link :to="{name: 'CareCoor'}">
            <button type="button" class="btn btn-secondary custom">Care Coordination</button>
        </router-link>
    </div>
</template>

<style scoped>
.custom{
    width: 150px;
    margin-bottom: 15px;
}
</style>