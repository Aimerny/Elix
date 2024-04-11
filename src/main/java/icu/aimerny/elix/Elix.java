package icu.aimerny.elix;

import icu.aimerny.elix.entity.ModEntities;
import icu.aimerny.elix.registry.*;
import net.fabricmc.api.ModInitializer;
import net.fabricmc.fabric.api.biome.v1.BiomeModifications;
import net.fabricmc.fabric.api.biome.v1.BiomeSelectors;
import net.minecraft.registry.Registry;
import net.minecraft.registry.RegistryKey;
import net.minecraft.registry.RegistryKeys;
import net.minecraft.util.Identifier;
import net.minecraft.world.gen.GenerationStep;
import net.minecraft.world.gen.feature.PlacedFeature;

public class Elix implements ModInitializer {


    public static final String MOD_ID = "elix";

    public static final RegistryKey<PlacedFeature> END_ROD_KEY = RegistryKey.of(RegistryKeys.PLACED_FEATURE, new Identifier(MOD_ID, "ore_custom"));

    @Override
    public void onInitialize() {
        ModItem.init();
        ModItemGroup.init();
        ModBlock.init();

        ModBlockEntity.init();
        ModEnchantment.init();
        ModEntities.init();

        BiomeModifications.addFeature(BiomeSelectors.foundInOverworld(), GenerationStep.Feature.UNDERGROUND_ORES, END_ROD_KEY);
    }
}
